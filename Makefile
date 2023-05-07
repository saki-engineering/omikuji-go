generate:
	cd internal && swig -go -c++ -cgo -includeall -I./include -intgosize 32 omikuji.i