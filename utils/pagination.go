package utils

// 获取分页的参数
func GetPage(pageNumber, pageSize int) (int, int) {
	var pageNumberResult, pageSizeResult int

	// 页码
	if pageSize != 0 {
		pageSizeResult = pageSize
	} else {
		pageSizeResult = 10
	}

	// 页数
	if pageNumber != 0 {
		pageNumberResult = (pageNumber - 1) * pageSizeResult
	} else {
		pageNumberResult = 1
	}

	return pageNumberResult, pageSizeResult
}
