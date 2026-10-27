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

	"(*go.uber.org/zap.Logger).Info":   0,
	"(*go.uber.org/zap.Logger).Error":  0,
	"(*go.uber.org/zap.Logger).Debug":  0,
	"(*go.uber.org/zap.Logger).Warn":   0,
	"(*go.uber.org/zap.Logger).DPanic": 0,
	"(*go.uber.org/zap.Logger).Fatal":  0,

	"(*go.uber.org/zap.SugaredLogger).Debug":  0,
	"(*go.uber.org/zap.SugaredLogger).Info":   0,
	"(*go.uber.org/zap.SugaredLogger).Warn":   0,
	"(*go.uber.org/zap.SugaredLogger).Error":  0,
	"(*go.uber.org/zap.SugaredLogger).DPanic": 0,
	"(*go.uber.org/zap.SugaredLogger).Panic":  0,
	"(*go.uber.org/zap.SugaredLogger).Fatal":  0,

	"(*go.uber.org/zap.SugaredLogger).Debugf":  0,
	"(*go.uber.org/zap.SugaredLogger).Infof":   0,
	"(*go.uber.org/zap.SugaredLogger).Warnf":   0,
	"(*go.uber.org/zap.SugaredLogger).Errorf":  0,
	"(*go.uber.org/zap.SugaredLogger).DPanicf": 0,
	"(*go.uber.org/zap.SugaredLogger).Panicf":  0,
	"(*go.uber.org/zap.SugaredLogger).Fatalf":  0,

	"(*go.uber.org/zap.SugaredLogger).Debugw":  0,
	"(*go.uber.org/zap.SugaredLogger).Infow":   0,
	"(*go.uber.org/zap.SugaredLogger).Warnw":   0,
	"(*go.uber.org/zap.SugaredLogger).Errorw":  0,
	"(*go.uber.org/zap.SugaredLogger).DPanicw": 0,
	"(*go.uber.org/zap.SugaredLogger).Panicw":  0,
	"(*go.uber.org/zap.SugaredLogger).Fatalw":  0,
}
