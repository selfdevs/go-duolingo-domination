package canvas

import (
	"duolingo/duolingo"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"strconv"
)

func drawTitle(ctx *canvas.Context) {
	ctx.DrawText(480, 1024, canvas.NewTextLine(createTitleFontFace(), "Duolingo Domination Leaderboard", canvas.Center))
}

func drawEntry(ctx *canvas.Context, user duolingo.User, index int) {
	y := 960 - float64(index)*55
	ctx.DrawText(40, y, canvas.NewTextLine(createEntryFontFace(), strconv.Itoa(index+1)+".", canvas.Left))
	if user.Streak > 0 {
		drawImage(ctx, "assets/fire.png", 120, y, 16)
	}
	ctx.DrawText(160, y, canvas.NewTextLine(createEntryFontFace(), strconv.Itoa(user.Streak), canvas.Left))
	ctx.DrawText(260, y, canvas.NewTextLine(createEntryFontFace(), user.Name, canvas.Left))
	ctx.DrawText(770, y, canvas.NewTextLine(createEntryFontFace(), strconv.Itoa(user.TotalXp)+" XP", canvas.Right))
	ctx.DrawText(790, y, canvas.NewTextLine(createEntryFontFace(), "+"+strconv.Itoa(user.GainedXp), canvas.Left))
}

func drawLogo(ctx *canvas.Context) {
	drawImage(ctx, "assets/logo.png", 425, 1088, 12)
	drawImage(ctx, "assets/target.png", 800, 1100, 0.75)
	drawImage(ctx, "assets/gem.png", 200, 1170, 0.75)
	drawImage(ctx, "assets/gem.png", 220, 1150, 0.75)
}

func drawBackground(ctx *canvas.Context) {
	ctx.SetFillColor(backgroundColor)
	ctx.DrawPath(0, 0, canvas.Rectangle(960, 1280))
}

func DrawLeaderboard(users []duolingo.User) {
	println("Starting leaderboard generation")
	C := canvas.New(960, 1280)
	ctx := canvas.NewContext(C)

	drawBackground(ctx)
	drawLogo(ctx)
	drawTitle(ctx)

	for i := 0; i < len(users); i++ {
		drawEntry(ctx, users[i], i)
	}

	err := renderers.Write("leaderboard.png", C, canvas.DPMM(1))
	if err != nil {
		println(err)
	}
	println("Done generating leaderboard, saved to leaderboard.png")
}
