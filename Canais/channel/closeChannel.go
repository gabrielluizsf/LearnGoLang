package channel

func closeINTchannel(send chan <- int){
  close(send);
}
