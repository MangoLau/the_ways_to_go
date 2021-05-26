package main


// 1 出于性能考虑建议使用带缓存的通道：
// 2 限制一个通道的数据数量并将它们封装成一个数组： ch := make(chan type,buf)