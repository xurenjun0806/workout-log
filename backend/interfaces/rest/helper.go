package rest

type ResponseError struct {
	Message string `json:"message"`
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s)) // 変換結果を記憶するスライス
	for i, v := range s {
		r[i] = f(v) // sの各要素を関数fで変換し結果をrに記憶
	}
	return r // 変換結果が記憶されたスライスを返す
}
