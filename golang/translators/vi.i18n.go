package translator

var (
	VietnameseTranslations = Translations{
		general_invalidId:          "ID không hợp lệ",
		general_createSuccess:      "Đã tạo xong",
		general_updateSuccess:      "Đã cập nhật",
		general_changeStateSuccess: "Trạng thái đã thay đổi",
		general_deleteSuccess:      "Đã xóa",
		general_error:              "Đã xảy ra lỗi",
		general_pleaseInputCorrect: "Vui lòng nhập thông tin chính xác",
		general_success:            "Thành công",

		// INFO MESSAGES
		infoMsg_signInSuccess:  "Đăng Nhập Thành Công",
		infoMsg_signOutSuccess: "Đăng Xuất Thành Công",
		infoMsg_createSuccess:  "Tạo thành công",
		infoMsg_updateSuccess:  "Cập nhật thành công",
		infoMsg_deleteSuccess:  "Xoá thành công",

		// ERROR MESSAGES
		errExceptionMsg_badRequest:           "Yêu Cầu Không Hợp Lệ",
		errExceptionMsg_unauthorized:         "Không Được Phép",
		errExceptionMsg_forbidden:            "Không Có Quyền",
		errExceptionMsg_unprocessableContent: "Nội Dung Không Thể Xử Lý",
		errExceptionMsg_internal:             "Lỗi Máy Chủ",
		errExceptionMsg_tournamentStarted:    "Trần đấu đã bắt đầu không thể thay đổi",

		// VALIDATION MESSAGES
		errValidationMsg_general:            "Đã có lỗi xảy ra, vui lòng kiểm tra lại dữ liệu bạn đã nhập",
		errValidationMsg_invalidEmailFormat: "Email không hợp lệ",
		errValidationMsg_required:           "bắt buộc",
		errValidationMsg_invalid:            "có giá trị ko hợp lệ",
		errValidation_wrongFormat:           "không hợp lệ",
		errValidation_minLength:             "có độ dài tối thiểu là %d",
		errValidation_maxLength:             "có độ dài tối đa là %d",
		errValidation_uniq:                  "Đã tồn tại",
		errValidation_notExist:              "Không tồn tại",
		errValidation_maxSizeImg:            "Kích thước hình ảnh tối đa có thể được tải lên là %dMB",
		errValidation_greaterThanInt:        "phải lớn hơn %d",
		errValidation_greaterThanOrEqualInt: "phải lớn hơn hoặc bằng %d",
		errValidation_lessThanOrEqualInt:    "phải nhỏ hơn hoặc bằng %d",
		errValidation_lessThanInt:           "phải nhỏ hơn %d",
		errValidation_greaterThanTime:       "phải sau %s",
		errValidation_lessThanTime:          "phải trước %s",
		errValidation_invalidJson:           "json không hợp lệ",
		errValidation_reachMaximum:          "Đã đạt đến giới hạn",
		errValidation_lessThanOrEqualFloat:  "phải nhỏ hơn hoặc bằng %.2f",
		errValidation_notDivisibleBy:        "phải chia hết cho %s",

		ValidationIsPowerOf: "phải là luỹ thừa của %d",

		errValidation_roleHierarchy:       "bạn không có quyền tạo tài khoản có role cao hơn",
		errValidation_roleUpdateHierarchy: "không thể cập nhật lên vị trí cao hơn",
		errValidation_roleDeleteHierarchy: "Không thể xoá người dùng có vị trí cao hơn",
		errValidation_deleteAdmin:         "Không thể xoá tài khoản admin",
		errValidation_updateHierarchy:     "Bạn không có quyền cập nhật thông tin tài khoản",
		errValidation_roleDelete:          "Bạn không có quyền xoá tài khoản người chơi trận đấu",

		errValidation_pagingOverflow: "Số trang đã vượt quá giới hạn",
		// DB Error Messages
		errDbMsg_notFound:      "Không tìm được",
		errDbMsg_unexpected:    "Lỗi DB không mong đợi",
		errDbMsg_wrongPassword: "Tên người dùng hoặc mật khẩu không chính xác",
	}
)
