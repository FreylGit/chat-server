package generate

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i github.com/FreylGit/platform_common/pkg/db.DB -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i github.com/FreylGit/platform_common/pkg/db.TxManager -o ./mocks/ -s "_minimock.go"
