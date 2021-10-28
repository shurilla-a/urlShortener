package logger

import "log"

func ErrorLogger(ErrorToFile error, MsgToFile string ) {
log.Panicf("%s:%s",ErrorToFile,MsgToFile )
}

//func InformToLog() {
//
//}