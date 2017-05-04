package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tarm/serial"
	"log"
	"fmt"
)

type MainController struct {
	beego.Controller

}

func (this *MainController) Get() {
	this.TplName = "index.tpl"

}

func (this *MainController) Up() {

	   c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        	
        _, err = s.Write([]byte("1"))
        if err != nil {
                log.Fatal(err)
        }
        	this.TplName = "index.tpl"
        fmt.Println(c)
}

func (this *MainController) Down() {

	   c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        _, err = s.Write([]byte("0"))
        if err != nil {
                log.Fatal(err)
        }
        this.TplName = "index.tpl"
}

func (this *MainController) Left() {

	   c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        _, err = s.Write([]byte("3"))
        if err != nil {
                log.Fatal(err)
        }

         this.TplName = "index.tpl"
}

func (this *MainController) Right() {

	   c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        _, err = s.Write([]byte("2"))
        if err != nil {
                log.Fatal(err)
        }

         this.TplName = "index.tpl"
}

func (this *MainController) Stop() {
      c := &serial.Config{Name: "/dev/ttyACM0", Baud:9600 }
      s, err := serial.OpenPort(c)

      if err != nil {
            log.Fatal(err)
      }

      _, err = s.Write([]byte("4"))

      if err != nil {
        log.Fatal(err)
      }

      this.TplName = "index.tpl"
}