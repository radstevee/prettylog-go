package prettylog

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/go-errors/errors"
)

type AnsiColor struct {
	Code string
}

const (
	RESET                    = "\033[0m"
	BOLD                     = "\033[1m"
	BLACK                    = "\033[30m"
	RED                      = "\033[31m"
	GREEN                    = "\033[32m"
	YELLOW                   = "\033[33m"
	BLUE                     = "\033[34m"
	PURPLE                   = "\033[35m"
	CYAN                     = "\033[36m"
	WHITE                    = "\033[37m"
	BRIGHT_BLACK             = "\033[90m"
	BRIGHT_RED               = "\033[91m"
	BRIGHT_GREEN             = "\033[92m"
	BRIGHT_YELLOW            = "\033[93m"
	BRIGHT_BLUE              = "\033[94m"
	BRIGHT_PURPLE            = "\033[95m"
	BRIGHT_CYAN              = "\033[96m"
	BRIGHT_WHITE             = "\033[97m"
	BLACK_BACKGROUND         = "\033[40m"
	RED_BACKGROUND           = "\033[41m"
	GREEN_BACKGROUND         = "\033[42m"
	YELLOW_BACKGROUND        = "\033[43m"
	BLUE_BACKGROUND          = "\033[44m"
	PURPLE_BACKGROUND        = "\033[45m"
	CYAN_BACKGROUND          = "\033[46m"
	WHITE_BACKGROUND         = "\033[47m"
	BRIGHT_BLACK_BACKGROUND  = "\033[100m"
	BRIGHT_RED_BACKGROUND    = "\033[101m"
	BRIGHT_GREEN_BACKGROUND  = "\033[102m"
	BRIGHT_YELLOW_BACKGROUND = "\033[103m"
	BRIGHT_BLUE_BACKGROUND   = "\033[104m"
	BRIGHT_PURPLE_BACKGROUND = "\033[105m"
	BRIGHT_CYAN_BACKGROUND   = "\033[106m"
	BRIGHT_WHITE_BACKGROUND  = "\033[107m"
)

func ForeColor(code int) string {
	return fmt.Sprintf("\033[38;5;%dm", code)
}

func BackColor(code int) string {
	return fmt.Sprintf("\033[48;5;%dm", code)
}

var (
	Gray                 = AnsiColor{ForeColor(244)}
	GrayBackground       = AnsiColor{BackColor(244)}
	Orange               = AnsiColor{ForeColor(202)}
	OrangeBackground     = AnsiColor{BackColor(202)}
	Pink                 = AnsiColor{ForeColor(200)}
	PinkBackground       = AnsiColor{BackColor(200)}
	CutePink             = AnsiColor{ForeColor(205)}
	CutePinkBackground   = AnsiColor{BackColor(205)}
	Aqua                 = AnsiColor{ForeColor(45)}
	AquaBackground       = AnsiColor{BackColor(45)}
	Gold                 = AnsiColor{ForeColor(220)}
	GoldBackground       = AnsiColor{BackColor(220)}
	LightGreen           = AnsiColor{ForeColor(82)}
	LightGreenBackground = AnsiColor{BackColor(82)}
	LightBlue            = AnsiColor{ForeColor(39)}
	LightBlueBackground  = AnsiColor{BackColor(39)}
	Magenta              = AnsiColor{ForeColor(13)}
	MagentaBackground    = AnsiColor{BackColor(13)}
	LightCyan            = AnsiColor{ForeColor(51)}
	LightCyanBackground  = AnsiColor{BackColor(51)}
	LightGray            = AnsiColor{ForeColor(247)}
	LightGrayBackground  = AnsiColor{BackColor(247)}
	DarkRed              = AnsiColor{ForeColor(88)}
	DarkRedBackground    = AnsiColor{BackColor(88)}
	DarkGreen            = AnsiColor{ForeColor(22)}
	DarkGreenBackground  = AnsiColor{BackColor(22)}
	DarkBlue             = AnsiColor{ForeColor(19)}
	DarkBlueBackground   = AnsiColor{BackColor(19)}
	DarkYellow           = AnsiColor{ForeColor(136)}
	DarkYellowBackground = AnsiColor{BackColor(136)}
	DarkPurple           = AnsiColor{ForeColor(55)}
	DarkPurpleBackground = AnsiColor{BackColor(55)}
)

type AnsiPair struct {
	Background AnsiColor
	Foreground AnsiColor
}

var (
	BlackPair        = AnsiPair{AnsiColor{BLACK_BACKGROUND}, AnsiColor{BLACK}}
	RedPair          = AnsiPair{AnsiColor{RED_BACKGROUND}, AnsiColor{RED}}
	GreenPair        = AnsiPair{AnsiColor{GREEN_BACKGROUND}, AnsiColor{GREEN}}
	YellowPair       = AnsiPair{AnsiColor{YELLOW_BACKGROUND}, AnsiColor{YELLOW}}
	BluePair         = AnsiPair{AnsiColor{BLUE_BACKGROUND}, AnsiColor{BLUE}}
	PurplePair       = AnsiPair{AnsiColor{PURPLE_BACKGROUND}, AnsiColor{PURPLE}}
	CyanPair         = AnsiPair{AnsiColor{CYAN_BACKGROUND}, AnsiColor{CYAN}}
	WhitePair        = AnsiPair{AnsiColor{WHITE_BACKGROUND}, AnsiColor{WHITE}}
	BrightBlackPair  = AnsiPair{AnsiColor{BRIGHT_BLACK_BACKGROUND}, AnsiColor{BRIGHT_BLACK}}
	BrightRedPair    = AnsiPair{AnsiColor{BRIGHT_RED_BACKGROUND}, AnsiColor{BRIGHT_RED}}
	BrightGreenPair  = AnsiPair{AnsiColor{BRIGHT_GREEN_BACKGROUND}, AnsiColor{BRIGHT_GREEN}}
	BrightYellowPair = AnsiPair{AnsiColor{BRIGHT_YELLOW_BACKGROUND}, AnsiColor{BRIGHT_YELLOW}}
	BrightBluePair   = AnsiPair{AnsiColor{BRIGHT_BLUE_BACKGROUND}, AnsiColor{BRIGHT_BLUE}}
	BrightPurplePair = AnsiPair{AnsiColor{BRIGHT_PURPLE_BACKGROUND}, AnsiColor{BRIGHT_PURPLE}}
	BrightCyanPair   = AnsiPair{AnsiColor{BRIGHT_CYAN_BACKGROUND}, AnsiColor{BRIGHT_CYAN}}
	BrightWhitePair  = AnsiPair{AnsiColor{BRIGHT_WHITE_BACKGROUND}, AnsiColor{BRIGHT_WHITE}}
	GrayPair         = AnsiPair{GrayBackground, Gray}
	OrangePair       = AnsiPair{OrangeBackground, Orange}
	PinkPair         = AnsiPair{PinkBackground, Pink}
	CutePinkPair     = AnsiPair{CutePinkBackground, CutePink}
	AquaPair         = AnsiPair{AquaBackground, Aqua}
	GoldPair         = AnsiPair{GoldBackground, Gold}
	LightGreenPair   = AnsiPair{LightGreenBackground, LightGreen}
	LightBluePair    = AnsiPair{LightBlueBackground, LightBlue}
	MagentaPair      = AnsiPair{MagentaBackground, Magenta}
	LightCyanPair    = AnsiPair{LightCyanBackground, LightCyan}
	LightGrayPair    = AnsiPair{LightGrayBackground, LightGray}
	DarkRedPair      = AnsiPair{DarkRedBackground, DarkRed}
	DarkGreenPair    = AnsiPair{DarkGreenBackground, DarkGreen}
	DarkBluePair     = AnsiPair{DarkBlueBackground, DarkBlue}
	DarkYellowPair   = AnsiPair{DarkYellowBackground, DarkYellow}
	DarkPurplePair   = AnsiPair{DarkPurpleBackground, DarkPurple}
)

type LogType struct {
	Emoji     string
	Name      string
	ColorPair AnsiPair
}

var (
	Information = LogType{"ℹ️ ", "Information", CyanPair}
	Runtime     = LogType{"✨", "Runtime", MagentaPair}
	Debug       = LogType{"🔧", "Debug", GrayPair}
	Network     = LogType{"🔌", "Network", BluePair}
	Success     = LogType{"✔️ ", "Success", BrightGreenPair}
	Warning     = LogType{"⚠️ ", "Warning", BrightYellowPair}
	Error       = LogType{"⛔", "Error", RedPair}
	Exception   = LogType{"💣", "Exception", RedPair}
	Critical    = LogType{"🚨", "Critical", BrightRedPair}
	Audit       = LogType{"📋", "Audit", YellowPair}
	Trace       = LogType{"🔍", "Trace", LightBluePair}
	Security    = LogType{"🔒", "Security", PurplePair}
	UserAction  = LogType{"🧍", "User Action", CutePinkPair}
	Performance = LogType{"⏱️ ", "Performance", PinkPair}
	Config      = LogType{"⚙️ ", "Config", LightGrayPair}
	Fatal       = LogType{"☠️ ", "Fatal", DarkRedPair}
)

type LoggerStyle string

const (
	FULL                      LoggerStyle = "<background><black><emoji> <prefix>: <message>"
	PREFIX                    LoggerStyle = "<background><black><emoji> <prefix>:<reset> <foreground><message>"
	SUFFIX                    LoggerStyle = "<foreground><emoji> <prefix>: <background><black><message>"
	TEXT_ONLY                 LoggerStyle = "<foreground><emoji> <prefix>: <message>"
	PREFIX_WHITE_TEXT         LoggerStyle = "<background><black><emoji> <prefix>:<reset> <message>"
	BRACKET_PREFIX            LoggerStyle = "<foreground><bold>[<emoji> <prefix>]<reset><foreground> <message>"
	BRACKET_PREFIX_WHITE_TEXT LoggerStyle = "<foreground><bold>[<emoji> <prefix>] <reset><message>"
)

var LoggerSettings = struct {
	LoggerStyle       LoggerStyle
	SaveToFile        bool
	SaveDirectoryPath string
	LogFileNameFormat string
}{
	LoggerStyle:       PREFIX,
	SaveToFile:        false,
	SaveDirectoryPath: "",
	LogFileNameFormat: "2006-01-02-150405",
}

var logFileName string = ""
var logFile string = ""

func InitLoggerFileWriter() {
	logFileName = time.Now().Format(LoggerSettings.LogFileNameFormat)
	if !strings.HasSuffix(LoggerSettings.SaveDirectoryPath, "/") {
		LoggerSettings.SaveDirectoryPath += "/"
	}

	logFile = LoggerSettings.SaveDirectoryPath + logFileName + ".log"

	_, err := os.Stat(logFile)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}

		os.MkdirAll(LoggerSettings.SaveDirectoryPath, fs.ModePerm)
		os.Create(logFile)
	}
}

func logToFile(message string, logType LogType) {
	time := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content := fmt.Sprintf("%s [%s] %s\n", time, strings.ToUpper(logType.Name), message)
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}

func Log(message string, logType LogType) {
	pattern := string(LoggerSettings.LoggerStyle)
	if logType == Fatal {
		pattern = string(FULL)
	}

	pattern = replacePlaceholders(pattern, logType, message)
	fmt.Println(pattern + RESET)

	if LoggerSettings.SaveToFile {
		logToFile(message, logType)
	}
}

func LogException(err error) {
	stacktrace := err.(*errors.Error).ErrorStack()
	scanner := bufio.NewScanner(strings.NewReader(stacktrace))

	for scanner.Scan() {
		Log(replaceAll(scanner.Text(), "\t", "  "), Exception)
	}
}

func replacePlaceholders(pattern string, logType LogType, message string) string {
	pattern = replaceAll(pattern, "<background>", logType.ColorPair.Background.Code)
	pattern = replaceAll(pattern, "<foreground>", logType.ColorPair.Foreground.Code)
	pattern = replaceAll(pattern, "<black>", BLACK)
	pattern = replaceAll(pattern, "<prefix>", logType.Name)
	pattern = replaceAll(pattern, "<message>", message)
	pattern = replaceAll(pattern, "<reset>", RESET)
	pattern = replaceAll(pattern, "<bold>", BOLD)
	pattern = replaceAll(pattern, "<emoji>", logType.Emoji)
	return pattern
}

func replaceAll(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
