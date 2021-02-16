package mp

import "fmt"

//接单接口
type Player interface {
	Play(source string)
}

//音乐播放模块应该是很容易扩展的，不应该在每次增加一种新音乐文件类型支持时都就需要大幅调整代码。
//简单但又足够通用的播放函数
//这里没有直接将MusicEntry作为参数传入，因为MusicEntry包含了一些多余的信息。
//本着最小原则，只需要将真正需要的信息传入即可，即音乐文件的位置以及音乐的类型
//入口
func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsuported music type", mtype)
		return
	}
	
	p.Play(source)
}
