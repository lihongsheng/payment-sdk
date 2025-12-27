package tools

import (
	"bytes"
	cRand "crypto/rand"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// 计算字符串的哈希值并映射到 [0, 1023] 范围
func GetPodNodeID() uint32 {
	serviceName := os.Getenv("HOSTNAME")
	if serviceName == "" {
		serviceName = GenerateRandomDigits(4)
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(serviceName))
	hash := h.Sum32()
	return hash % 1024
}

func GenerateRandomDigits(length uint32) string {
	l := int(length)
	// 计算上限值，例如长度为 3 时，上限为 1000
	maxBig := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(l)), nil)
	for {
		// 生成随机数
		num, err := cRand.Int(cRand.Reader, maxBig)
		if err != nil {
			return GenerateRandomUnix(l)
		}
		// 转换为字符串
		result := num.String()
		// 若长度不足，前面补 0
		for len(result) < l {
			result = "0" + result
		}
		return result
	}
}

// GenerateRandomString 生成指定长度的随机字符串，包含0-9和a-z
func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}

	// 初始化随机数生成器，使用当前时间作为种子
	rand.Seed(time.Now().UnixNano())

	// 定义字符集：0-9和a-z
	charset := "0123456789abcdefghijklmnopqrstuvwxyz"
	charsetLen := len(charset)

	// 创建结果切片
	result := make([]byte, length)

	// 随机选择字符
	for i := 0; i < length; i++ {
		// 生成0到charsetLen-1之间的随机索引
		index := rand.Intn(charsetLen)
		result[i] = charset[index]
	}

	return string(result)
}

func GenerateRandomUnix(length int) string {
	if length <= 0 {
		return ""
	}
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		if i == 0 {
			// 首位不能为 0
			digit := rand.Intn(9) + 1
			result += strconv.Itoa(digit)
		} else {
			digit := rand.Intn(10)
			result += strconv.Itoa(digit)
		}
	}
	return result
}

var once sync.Once
var generateRandom *GenerateRandom

func GetID() string {
	once.Do(func() {
		node := GetPodNodeID()
		generateRandom = NewGenerate(node)
	})
	return generateRandom.Generate()
}

// GenerateRandom
// 13 + 4 + 4 + 1 + 4 = 26 长度
// 时钟回滚的时候，标志位1，正常下为0。 时钟回滚也允许利用当前回滚时间继续生成，用标志位0进行判断
// +-----------------------------------------------------------------------------------------+
// | 时间戳:1757061730927 | 机器码:0001-1024 | 自增id:0001-4096 | 标志位：0 时钟未回滚，1时钟回滚 | 随机数:0001-9999
// +-----------------------------------------------------------------------------------------+
type GenerateRandom struct {
	node          string
	step          int64
	stepMax       int64
	lastTimestamp int64
	preTimestamp  int64
	mutex         sync.Mutex
}

func NewGenerate(node uint32) *GenerateRandom {
	nodeStr := fmt.Sprintf("%04d", node)
	return &GenerateRandom{
		node:    nodeStr,
		step:    0,
		stepMax: 4096,
		mutex:   sync.Mutex{},
	}
}

// Generate
// 时钟回滚的时候，标志位1，正常下为0。 时钟回滚也允许利用当前回滚时间继续生成，用标志位0进行判断
func (s *GenerateRandom) Generate() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	mode := 0
	currentTime := time.Now().UnixMilli()
	if currentTime < s.lastTimestamp { // 时钟回滚
		s.step = (s.step + 1) & s.stepMax
		if s.step == 0 {
			currentTime = s.timeMills(s.lastTimestamp)
		}
		if currentTime < s.lastTimestamp {
			currentTime = s.lastTimestamp
		}
		mode = 1
	} else if currentTime == s.lastTimestamp {
		s.step = (s.step + 1) & s.stepMax
		if s.step == 0 {
			currentTime = s.timeMills(s.lastTimestamp)
		}
	} else {
		s.step = 0
	}
	s.lastTimestamp = currentTime
	step := fmt.Sprintf("%04d", s.step)
	return fmt.Sprintf("%d%s%s%d%s", currentTime, s.node, step, mode, GenerateRandomDigits(4))
}

// 等待下一毫秒
func (s *GenerateRandom) timeMills(lastTime int64) int64 {
	currentTime := time.Now().UnixMilli()
	for currentTime <= lastTime {
		currentTime = time.Now().UnixMilli()
	}
	return currentTime
}

func GetRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

func GetResponseBody(request *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}
	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
