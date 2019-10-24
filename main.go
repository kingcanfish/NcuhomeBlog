package main

import "NcuhomeBlog/views"

func main()  {
	app := views.GetRoute()
	app.Run(":8087")

}
