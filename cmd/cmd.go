package cmd

import (
	"bufio"
	"fmt"
	"github.com/Swmiao1/tarstools/util"
	"io"
	"os"
	"os/exec"
	"strings"
)

type cmdText struct {
	Text string
	err  bool
}

type GTPConnection struct {
	Cmd     *exec.Cmd
	infile  io.WriteCloser
	outfile io.ReadCloser
	errfile io.ReadCloser
	outText chan cmdText
	IsErr   bool
}

func (c *GTPConnection) Input(s string) {
	//判断是否有 /n 结尾
	temp := strings.Split(s, "")
	l := len(temp)
	if l > 1 && temp[l-1] != "\n" {
		s += "\n"
	}
	if s == "\n" {
		return
	}
	n, err := c.infile.Write([]byte(fmt.Sprintf("%s", s)))
	if err != nil {
		fmt.Println("[IN]", n, s, err)
	}
}

func (c GTPConnection) Listen() {
	go c.listenOut()
	go c.listenErr()
}

//监听返回文本
func (c *GTPConnection) listenOut() {
	reader := bufio.NewReader(c.outfile)
	for {
		var buf = make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil || io.EOF == err {
			break
		}
		str := string(util.ConvertToByte(string(buf[:n]), "gbk", "utf8"))
		if str != "\r\n" {
			c.outText <- cmdText{Text: str}
		}

	}
}

//监听错误返回文本
func (c *GTPConnection) listenErr() {
	reader := bufio.NewReader(c.errfile)
	for {
		var buf = make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil || io.EOF == err {
			break
		}
		str := string(util.ConvertToByte(string(buf[:n]), "gbk", "utf8"))
		fmt.Fprintln(os.Stderr, str)

		c.IsErr = true
		//c.outText <- cmdText{Text: str, err: true}
		os.Exit(0)
	}
}

//打印结果
func (c GTPConnection) Print() {
	text := <-c.outText
	fmt.Print(text.Text)
}
func (c GTPConnection) Log() {
	go func() {
		for {
			text := <-c.outText
			//if text.Text[0] == '\n' {
			//	text.Text = text.Text[2:]
			//}
			fmt.Print(text.Text)
		}
	}()

}

func NewCmd() *GTPConnection {
	cmd := GTPConnection{
		Cmd:     exec.Command("cmd", "run"),
		outText: make(chan cmdText, 1024),
	}
	var err error
	//创建输入管道
	cmd.infile, err = cmd.Cmd.StdinPipe()
	//创建输出管道
	cmd.outfile, err = cmd.Cmd.StdoutPipe()
	//创建输出管道
	cmd.errfile, err = cmd.Cmd.StderrPipe()
	err = cmd.Cmd.Start()
	if err != nil {
		println(err)
	}
	//监听输出
	cmd.Listen()
	return &cmd
}
