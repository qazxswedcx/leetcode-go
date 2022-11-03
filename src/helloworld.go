package main

//func main() {
//	//新建一张图表
//	p := plot.New()
//	//设置图表标题为空
//	p.Title.Text = "函数曲线"
//	//隐藏坐标轴
//	// p.HideAxes()
//	//为第一条（Tanh函数）曲线分配1000点容量的空间
//	points1 := make(plotter.XYs, 0, 1000)
//	//循环获取每个x点的Tanh函数值
//	for x := -10.0; x <= 10.0; x = x + 0.1 {
//		y := math.Cos(x)
//		points1 = append(points1, plotter.XY{X: x, Y: y})
//	}
//	//新建一条折线
//	line1, _ := plotter.NewLine(points1)
//	//设置画线使用的颜色、宽度、线型等
//	line1.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 0xFF, A: 0xFF}
//	line1.LineStyle.Width = vg.Points(2)
//	line1.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(6)}
//	//在图表中加入该折线
//	p.Add(line1)
//	//保存图表到图片文件
//	err := p.Save(8*vg.Inch, 8*vg.Inch, "points.png")
//	if err != nil {
//		return
//	}
//	fmt.Println("over")
//}
