package analyzer

// messageArgIndex сопоставляет полное имя функции логгера (как его возвращает
// (*types.Func).FullName()) с индексом аргумента, в котором лежит сообщение.
// Функции, которых нет в этой таблице, анализатор не проверяет.
var messageArgIndex = map[string]int{
	// log: пакетные функции
	"log.Print":   0,
	"log.Printf":  0,
	"log.Println": 0,
	"log.Fatal":   0,
	"log.Fatalf":  0,
	"log.Fatalln": 0,
	"log.Panic":   0,
	"log.Panicf":  0,
	"log.Panicln": 0,

	// log: методы *log.Logger
	"(*log.Logger).Print":   0,
	"(*log.Logger).Printf":  0,
	"(*log.Logger).Println": 0,
	"(*log.Logger).Fatal":   0,
	"(*log.Logger).Fatalf":  0,
	"(*log.Logger).Fatalln": 0,
	"(*log.Logger).Panic":   0,
	"(*log.Logger).Panicf":  0,
	"(*log.Logger).Panicln": 0,

	// log/slog: пакетные функции
	"log/slog.Debug":        0,
	"log/slog.Info":         0,
	"log/slog.Warn":         0,
	"log/slog.Error":        0,
	"log/slog.DebugContext": 1, // (ctx, msg, ...)
	"log/slog.InfoContext":  1,
	"log/slog.WarnContext":  1,
	"log/slog.ErrorContext": 1,
	"log/slog.Log":          2, // (ctx, level, msg, ...)
	"log/slog.LogAttrs":     2,

	// log/slog: методы *slog.Logger
	"(*log/slog.Logger).Debug":        0,
	"(*log/slog.Logger).Info":         0,
	"(*log/slog.Logger).Warn":         0,
	"(*log/slog.Logger).Error":        0,
	"(*log/slog.Logger).DebugContext": 1,
	"(*log/slog.Logger).InfoContext":  1,
	"(*log/slog.Logger).WarnContext":  1,
	"(*log/slog.Logger).ErrorContext": 1,
	"(*log/slog.Logger).Log":          2,
	"(*log/slog.Logger).LogAttrs":     2,
}
