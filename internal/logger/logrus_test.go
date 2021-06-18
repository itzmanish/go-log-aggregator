package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func newTestLogger() (Logger, *test.Hook) {
	logger, hook := test.NewNullLogger()
	l := &logrusLogger{
		logrus: logger,
	}
	l.Init(WithLevel(DebugLevel), WithOutput(os.Stderr))
	return l, hook
}

func TestLogger(t *testing.T) {
	logger, hook := newTestLogger()
	t.Run("test logger", func(t *testing.T) {
		assert.Equal(t, logger.String(), "Logrus logger")
	})
	t.Run("test error", func(t *testing.T) {
		logger.Error("Helloerror")
		logger.Errorf("%s", "Helloerror")
		assert.Equal(t, 2, len(hook.Entries))
		assert.Equal(t, logrus.Level(ErrorLevel), hook.LastEntry().Level)
		assert.Equal(t, "Helloerror", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})
	t.Run("test Info", func(t *testing.T) {
		logger.Info("Helloerror")
		logger.Infof("%s", "Helloerror")
		assert.Equal(t, 2, len(hook.Entries))
		assert.Equal(t, logrus.Level(InfoLevel), hook.LastEntry().Level)
		assert.Equal(t, "Helloerror", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})
	t.Run("test Debug", func(t *testing.T) {
		logger.Debug("Helloerror")
		logger.Debugf("%s", "Helloerror")
		assert.Equal(t, 2, len(hook.Entries))
		assert.Equal(t, logrus.Level(DebugLevel), hook.LastEntry().Level)
		assert.Equal(t, "Helloerror", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})
	t.Run("test Warn", func(t *testing.T) {
		logger.Warn("Helloerror")
		logger.Warnf("%s", "Helloerror")
		assert.Equal(t, 2, len(hook.Entries))
		assert.Equal(t, logrus.Level(WarnLevel), hook.LastEntry().Level)
		assert.Equal(t, "Helloerror", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("test Log", func(t *testing.T) {
		logger.Log(InfoLevel, "Helloerror")
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.Level(InfoLevel), hook.LastEntry().Level)
		assert.Equal(t, "Helloerror", hook.LastEntry().Message)

		hook.Reset()
		assert.Nil(t, hook.LastEntry())
	})

}
