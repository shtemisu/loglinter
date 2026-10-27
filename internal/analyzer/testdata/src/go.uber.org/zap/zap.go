// Package zap — заглушка go.uber.org/zap для analysistest.
// Повторяет только сигнатуры настоящего пакета; тела пустые.
package zap

type Field struct{}

func Error(err error) Field         { return Field{} }
func String(key, val string) Field  { return Field{} }
func Int(key string, val int) Field { return Field{} }
func Any(key string, val any) Field { return Field{} }

type Logger struct{}

func NewNop() *Logger { return &Logger{} }

func (l *Logger) Sugar() *SugaredLogger { return &SugaredLogger{} }

func (l *Logger) Debug(msg string, fields ...Field)  {}
func (l *Logger) Info(msg string, fields ...Field)   {}
func (l *Logger) Warn(msg string, fields ...Field)   {}
func (l *Logger) Error(msg string, fields ...Field)  {}
func (l *Logger) DPanic(msg string, fields ...Field) {}
func (l *Logger) Panic(msg string, fields ...Field)  {}
func (l *Logger) Fatal(msg string, fields ...Field)  {}

type SugaredLogger struct{}

func (s *SugaredLogger) Debug(args ...any)  {}
func (s *SugaredLogger) Info(args ...any)   {}
func (s *SugaredLogger) Warn(args ...any)   {}
func (s *SugaredLogger) Error(args ...any)  {}
func (s *SugaredLogger) DPanic(args ...any) {}
func (s *SugaredLogger) Panic(args ...any)  {}
func (s *SugaredLogger) Fatal(args ...any)  {}

func (s *SugaredLogger) Debugf(template string, args ...any)  {}
func (s *SugaredLogger) Infof(template string, args ...any)   {}
func (s *SugaredLogger) Warnf(template string, args ...any)   {}
func (s *SugaredLogger) Errorf(template string, args ...any)  {}
func (s *SugaredLogger) DPanicf(template string, args ...any) {}
func (s *SugaredLogger) Panicf(template string, args ...any)  {}
func (s *SugaredLogger) Fatalf(template string, args ...any)  {}

func (s *SugaredLogger) Debugw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Infow(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...any) {}
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...any)  {}
