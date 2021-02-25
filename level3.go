package main

import "fmt"

//现在需要一个表示视频详情的结构体（请自行设计，可以参考哔哩哔哩的），并且能实现
//点赞
//收藏
//投币
//一键三连
//这几个方法。
//而且还需实现一个发布视频的函数，入参为作者名，视频名，返回一个视频结构体

type Author struct {
	Name string //昵称
	VIP bool
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}

type Video struct {
	Title string
	ID string
	Author string
	viewCounts int //播放量
	dianzanCounts int //点赞数
	shoucangCounts int //收藏数
	toubiCounts int //投币数
}

type Comment struct {
	Author
	words string
	dianzanCounts int //点赞数
}

type videoDetial struct {
	Author
	Video
	Comment
}

func (d *videoDetial) dianzan()  {
	fmt.Println("点赞成功")
	d.Video.dianzanCounts++
}
func (d *videoDetial) toubi()  {
	fmt.Println("投币成功")
	d.Video.toubiCounts++
}
func (d *videoDetial) shoucang()  {
	fmt.Println("收藏成功")
	d.Video.shoucangCounts++
}
func (d *videoDetial) sanlian()  {
	fmt.Println("一键三连成功")
	d.Video.dianzanCounts++
	d.Video.toubiCounts++
	d.Video.shoucangCounts++
}

func faBuVideo(author string,videoTitle string) Video {
	v := Video{Author: author,Title: videoTitle}
	return v
}

func main()  {
	videoDetial := &videoDetial{}

	videoDetial.dianzan()
	fmt.Println("点赞数：",videoDetial.Video.dianzanCounts)
	videoDetial.toubi()
	fmt.Println("投币数：",videoDetial.Video.toubiCounts)
	videoDetial.shoucang()
	fmt.Println("收藏数：",videoDetial.Video.shoucangCounts)
	videoDetial.sanlian()
	fmt.Println("点赞数：",videoDetial.Video.dianzanCounts)
	fmt.Println("投币数：",videoDetial.Video.toubiCounts)
	fmt.Println("收藏数：",videoDetial.Video.shoucangCounts)
}
