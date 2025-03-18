package translator

var (
	EnglishTranslations = Translations{
		general_invalidId:          "Invalid ID",
		general_createSuccess:      "Created",
		general_updateSuccess:      "Updated",
		general_changeStateSuccess: "State Changed",
		general_deleteSuccess:      "Deleted",
		general_error:              "An error has occurred",
		general_pleaseInputCorrect: "Please enter the correct input",
		general_success:            "Success",

		// TOURNAMENT STATE STATUS
		tournamentState_publish:   "Published",
		tournamentState_unpublish: "Unpublish",

		// INFO MESSAGES
		infoMsg_signInSuccess:  "Sign In Successfully",
		infoMsg_signOutSuccess: "Sign Out Successfully",
		infoMsg_createSuccess:  "Create Successfully",
		infoMsg_updateSuccess:  "Update Successfully",
		infoMsg_deleteSuccess:  "Delete Successfully",

		// Exception ERROR MESSAGES
		errExceptionMsg_badRequest:           "Bad Request",
		errExceptionMsg_unauthorized:         "Unauthorized",
		errExceptionMsg_forbidden:            "Forbidden",
		errExceptionMsg_unprocessableContent: "Unprocessable Content",
		errExceptionMsg_internal:             "Internal Server Error",
		errExceptionMsg_tournamentStarted:    "Tournament has started do not update",

		// VALIDATION MESSAGES
		errValidationMsg_general:            "Errors Happened, Please check your input",
		errValidationMsg_required:           "is required",
		errValidationMsg_invalid:            "has invalid value",
		errValidation_wrongFormat:           "is wrong format",
		errValidation_minLength:             "is too short %d",
		errValidation_maxLength:             "is too long %d",
		errValidation_uniq:                  "already exists",
		errValidation_notExist:              "is not exist",
		errValidation_maxSizeImg:            "The maximum image size that can be uploaded is %dMB",
		errValidation_greaterThanInt:        "must be greater than %d",
		errValidation_greaterThanOrEqualInt: "must be greater than or equal to %d",
		errValidation_lessThanOrEqualInt:    "must be less than or equal to %d",
		errValidation_lessThanInt:           "must be less than %d",
		errValidation_greaterThanTime:       "must be after %s",
		errValidation_lessThanTime:          "must be before %s",
		errValidation_invalidJson:           "invalid json",
		errValidation_reachMaximum:          "Reached maximum",
		errValidation_lessThanOrEqualFloat:  "must be less than or equal to %.2f",
		errValidation_notDivisibleBy:        "must be divisible by %s",

		ValidationIsPowerOf: "must be power of %d",

		errValidation_roleHierarchy:       "You do not have permission to create an club user with a higher role.",
		errValidation_roleUpdateHierarchy: "can not update to higher role",
		errValidation_roleDeleteHierarchy: "Can not delete higher role user",
		errValidation_deleteAdmin:         "Can not delete admin account",
		errValidation_updateHierarchy:     "You do not have permission to update account information",
		errValidation_roleDelete:          "You do not have permission to delete an tournament player.",
		errValidation_pagingOverflow:      "Number of pages is overflow",
		// DB Error Messages
		errDbMsg_notFound:      "Not Found",
		errDbMsg_unexpected:    "Unexpected DB Error",
		errDbMsg_wrongPassword: "Wrong Username Or Password",
	}
)
