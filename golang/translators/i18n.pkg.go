package translator

type Translations struct {
	general_invalidId          string
	general_createSuccess      string
	general_deleteSuccess      string
	general_updateSuccess      string
	general_changeStateSuccess string
	general_error              string
	general_pleaseInputCorrect string
	general_success            string

	clubUserRoleType_admin   string
	clubUserRoleType_manager string
	clubUserRoleType_staff   string

	tournamentPlayerStatus_pending  string
	tournamentPlayerStatus_approved string
	tournamentPlayerStatus_denied   string

	// TOURNAMENT GAME TYPE
	tournamentGameType_10_ball   string
	tournamentGameType_9_ball    string
	tournamentGameType_8_ball    string
	tournamentGameType_3_cushion string

	// TOURNAMENT PLAYER TYPE
	tournamentPlayerType_single string
	tournamentPlayerType_double string
	tournamentPlayerType_teams  string

	// TOURNAMENTSTAGE STAGE TYPE
	tournamentStageStageType_single      string
	tournamentStageStageType_double      string
	tournamentStageStageType_round_robin string
	tournamentStageStageType_swiss       string

	// TOURNAMENTSTAGE STAGE TYPE
	tournamentBracketOrdering_random string
	tournamentBracketOrdering_manual string
	tournamentBracketOrdering_seed   string

	// TOURNAMENTSTAGE STATE TYPE
	tournamentState_publish   string
	tournamentState_unpublish string

	// INFO MESSAGES
	infoMsg_signInSuccess  string
	infoMsg_signOutSuccess string
	infoMsg_createSuccess  string
	infoMsg_updateSuccess  string
	infoMsg_deleteSuccess  string

	// EXCEPTION ERROR MESSAGES
	errExceptionMsg_badRequest           string
	errExceptionMsg_unauthorized         string
	errExceptionMsg_forbidden            string
	errExceptionMsg_unprocessableContent string
	errExceptionMsg_internal             string
	errExceptionMsg_tournamentStarted    string

	// Validation MESSAGES
	errValidationMsg_general            string
	errValidationMsg_invalidEmailFormat string
	errValidationMsg_required           string
	errValidationMsg_invalid            string
	errValidation_wrongFormat           string
	errValidation_minLength             string
	errValidation_maxLength             string
	errValidation_uniq                  string
	errValidation_notExist              string
	errValidation_maxSizeImg            string

	errValidation_greaterThanInt        string
	errValidation_greaterThanOrEqualInt string
	errValidation_lessThanOrEqualInt    string
	errValidation_lessThanInt           string
	errValidation_greaterThanTime       string
	errValidation_lessThanTime          string
	errValidation_invalidJson           string
	errValidation_reachMaximum          string
	errValidation_lessThanOrEqualFloat  string
	errValidation_notDivisibleBy        string

	ValidationIsPowerOf string

	errValidation_roleHierarchy       string
	errValidation_roleUpdateHierarchy string
	errValidation_roleDeleteHierarchy string
	errValidation_deleteAdmin         string
	errValidation_updateHierarchy     string
	errValidation_roleDelete          string
	errValidation_pagingOverflow      string
	// DB Error Messages
	errDbMsg_notFound      string
	errDbMsg_unexpected    string
	errDbMsg_wrongPassword string

	// Tournament table
	Title         string
	LiveStreamUrl string
	StartNumber   string
	EndNumber     string

	// Tournament player
	TournamentPlayersFull             string
	TournamentPlayerAlreadyRegistered string

	// Tournament Manager
	ErrTournamentManagerExists string
	ExternalClubUserIDs        string
}

var supportedLanguages = map[string]Translations{
	"en": EnglishTranslations,
	"vi": VietnameseTranslations,
}
