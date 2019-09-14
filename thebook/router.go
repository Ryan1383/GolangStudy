type router struct {
	// 키 : http 메서드
	// 값 : URL 패턴별로 실행할 HandlerFunc
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인

	m, ok := r.handlers[method]

	if ok == false {
		//등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]http.Handlefunc)
		r.handlers[method] = m
	}
	//http 메소드로 등록된 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}