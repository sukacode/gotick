stable locator
stable session
retry mechanism
logging


clean architecture Go
full folder structure
main.go
browser manager
session manager
notifier
scheduler
Dockerfile
config env
logging


> $go mod init gotick/hello




mkdir -p \
cmd/app \
internal/config \
internal/browser \
internal/auth \
internal/event \
internal/checkout \
internal/notifier \
internal/scheduler \
internal/logger \
internal/utils \
sessions \
screenshots \
logs

touch \
cmd/app/main.go \
internal/config/config.go \
internal/browser/playwright.go \
internal/browser/session.go \
internal/auth/login.go \
internal/event/monitor.go \
internal/event/locator.go \
internal/event/watcher.go \
internal/checkout/quantity.go \
internal/checkout/agreement.go \
internal/checkout/submit.go \
internal/notifier/telegram.go \
internal/scheduler/cron.go \
internal/logger/logger.go \
internal/utils/retry.go \
internal/utils/screenshot.go \
.env \
Dockerfile




go get github.com/playwright-community/playwright-go

go run github.com/playwright-community/playwright-go/cmd/playwright install