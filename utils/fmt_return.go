package utils

func FmtErrorReturn(err error) (map[string]interface{}, error) {
	return map[string]interface{}{
		"code":    1,
		"message": err.Error(),
		"data":    "",
	}, err
}

func FmtNormalReturn(data interface{}, msg ...string) (map[string]interface{}, error) {
	if len(msg) != 0 {
		return map[string]interface{}{
			"code":    0,
			"message": msg[0],
			"data":    data,
		}, nil
	} else {
		return map[string]interface{}{
			"code":    0,
			"message": "获取成功",
			"data":    data,
		}, nil
	}

}

