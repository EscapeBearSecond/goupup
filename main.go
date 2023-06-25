package main

import (
	file "github.com/study/golang/File/case"
)

func main() {
	// generic.FunctionCase()
	// generic.ClassCase()
	// generic.InterfaceCase()
	// Library.EncodingCase()
	// Library.ErrorsCase()
	// Library.RegexpCase()
	// Library.SortCase()
	// _defer.DeferCase1()
	// _defer.FileCase()
	// f1 := function.Sum
	// f2 := function.LogMiddleWare(f1)
	// f3 := function.LogMiddleWare(f2)
	// fmt.Println(f3(1, 2))
	// f4 := function.MyMiddleFun(f1)
	// fmt.Println(f4.Accumulation(1, 2, 3, 4))
	// fmt.Println(f2.Accumulation(1, 2, 3, 4))
	// fmt.Println(f3.Accumulation(1, 2, 3, 4))
	// fmt.Println(function.Fibonacci(10))
	// function.ClousreTrap()
	// function.ClosureSucc()

	// channel
	// channel.NoticeAndSelect()

	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	// defer cancel()
	// <- ctx.Done()

	// file case
	// file.CopyDirToDir()
	// file.ReadWriteFile()
	// file.RwSync()
	// file.Readline1()
	// file.Readline2()
	file.Readline3()
}
