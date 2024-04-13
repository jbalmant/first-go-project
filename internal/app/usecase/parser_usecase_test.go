package usecase_test

import (
	"testing"

	"MyFirstModule/internal/app/entity"
	"MyFirstModule/internal/app/usecase"
	"MyFirstModule/test_utils"
)

func TestMapLogToAddUserEventShouldPanicWhenPatternDoesNotMatch(t *testing.T) {
	parser := usecase.NewQuake3Parser()

	result := parser.Parse("WrongEventPatternSample")

	if result != nil {
		t.Errorf("Parser should return null")
	}
}

func TestShouldMapToAddUserEvent(t *testing.T) {
	testCases := []struct {
		input    string
		expected entity.Event
	}{
		{` 21:51 ClientUserinfoChanged: 3 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0`, *entity.NewAddPlayerEvent("Dono da Bola")},
		{` 21:53 ClientUserinfoChanged: 3 n\Mocinha\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0`, *entity.NewAddPlayerEvent("Mocinha")},
	}

	for _, tc := range testCases {
		// Arrange
		parser := usecase.NewQuake3Parser()

		// Act
		result := parser.Parse(tc.input)

		// Assert
		currentEvent := result.Details.(*entity.AddPlayerEvent)
		expectedEvent := tc.expected.Details.(*entity.AddPlayerEvent)
		test_utils.AssertObjectEquals(t, expectedEvent, currentEvent)
	}
}

func TestShouldMapToKillEvent(t *testing.T) {
	testCases := []struct {
		input    string
		expected entity.Event
	}{
		{" 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT", *entity.NewKillEvent("<world>", "Isgalamido", "MOD_TRIGGER_HURT")},
		{"  3:37 Kill: 3 4 7: Isgalamido killed Zeh by MOD_ROCKET_SPLASH", *entity.NewKillEvent("Isgalamido", "Zeh", "MOD_ROCKET_SPLASH")},
		{"  3:41 Kill: 2 3 6: Dono da Bola killed Isgalamido by MOD_ROCKET", *entity.NewKillEvent("Dono da Bola", "Isgalamido", "MOD_ROCKET")},
		{"123:41 Kill: 2 3 6: Dono da Bola killed Joãozinho da Silva Santos by MOD_ROCKET", *entity.NewKillEvent("Dono da Bola", "Joãozinho da Silva Santos", "MOD_ROCKET")},
	}

	for _, tc := range testCases {
		// Arrange
		parser := usecase.NewQuake3Parser()

		// Act
		result := parser.Parse(tc.input)

		// Assert
		test_utils.AssertObjectEquals(t, result.Details.(*entity.KillEvent), tc.expected.Details.(*entity.KillEvent))
	}
}

func TestShouldParseInitGameEvent(t *testing.T) {
	// Arrange
	line := `  0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`
	parser := usecase.NewQuake3Parser()
	expectedType := entity.Init

	// Act
	result := parser.Parse(line)

	// Assert
	if expectedType != result.Type {
		t.Errorf("Expected %q, got %q", expectedType, result.Type)
	}
}

func TestShouldParseEndGameEvent(t *testing.T) {
	// Arrange
	line := ` 20:37 ShutdownGame:`
	parser := usecase.NewQuake3Parser()
	expectedType := entity.End

	// Act
	result := parser.Parse(line)

	// Assert
	if expectedType != result.Type {
		t.Errorf("Expected %q, got %q", expectedType, result.Type)
	}
}
