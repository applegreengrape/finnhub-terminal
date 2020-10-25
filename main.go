package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/barchart"

	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/applegreengrape/finnhub-terminal/widgets"
	"github.com/applegreengrape/finnhub-terminal/yahoo"
)

func main() {
	// new terminal
	t, err := termbox.New()
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	// new ctx sessions
	ctx, cancel := context.WithCancel(context.Background())

	// update time
	timeNow, err := text.New()
	if err != nil {
		log.Fatal(err)
	}
	go widgets.UpdateTime(ctx, timeNow)

	// update stocks
	stk, err := text.New()
	if err != nil {
		log.Fatal(err)
	}
	go widgets.UpdateStockPrice(ctx, stk)

	// stock line chart
	cfg := yahoo.NewSettingFromConfig()
	const redrawInterval = 250 * time.Millisecond
	lc, err := linechart.New(
		linechart.YAxisAdaptive(),
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
	)
	if err != nil {
		log.Fatal(err)
	}
	go widgets.StockLineChart(ctx, lc, redrawInterval/3)

	// volume spark chart
	volchart, err := sparkline.New(
		sparkline.Label(fmt.Sprintf("%s volume chart", cfg.Stocks[0]), cell.FgColor(cell.ColorGreen)),
		sparkline.Color(cell.ColorYellow),
	)
	if err != nil {
		log.Fatal(err)
	}
	go widgets.VolumeSparkChart(ctx, volchart, redrawInterval)

	//RollingEarningsCalendar Widgest
	earningCals, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		log.Fatal(err)
	}
	go widgets.RollingEarningsCalendar(ctx, earningCals)

	// RollingMtkNews Widgest
	news, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		log.Fatal(err)
	}
	go widgets.RollingNews(ctx, news)

	//RollingCompanyNews Widgest
	companyNewsRoll, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		log.Fatal(err)
	}
	go widgets.RollingCompanyNews(ctx, companyNewsRoll)

	// update basic financials
	basicFinancials, err := text.New()
	if err != nil {
		log.Fatal(err)
	}
	go widgets.UpdateBasicFinancials(ctx, basicFinancials)

	// export all basic financials to csv button
	triggers := make(chan bool)
	go data.ExportAllBFs(triggers)
	export1, err := button.New(
		"Export all to csv",
		func() error {
			triggers <- true
			return nil
		},
		button.GlobalKey('e'),
		button.FillColor(cell.ColorNumber(220)),
		button.WidthFor("Export all to csv"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// export all financial reports to csv button
	triggers2 := make(chan bool)
	go data.ExportAllFRs(triggers2)
	export2, err := button.New(
		"Export all to csv",
		func() error {
			triggers2 <- true
			return nil
		},
		button.GlobalKey('e'),
		button.FillColor(cell.ColorNumber(220)),
		button.WidthFor("Export all to csv"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// update financial reports
	financialReports, err := text.New()
	if err != nil {
		log.Fatal(err)
	}
	go widgets.UpdateFinancialReports(ctx, financialReports)
	

	// target bar chart 
	bcTarget, err := barchart.New(
		barchart.BarColors([]cell.Color{
			cell.ColorGreen,
			cell.ColorGreen,
			cell.ColorYellow,
			cell.ColorRed,
			cell.ColorRed,
		}),
		barchart.ValueColors([]cell.Color{
			cell.ColorBlack,
			cell.ColorBlack,
			cell.ColorBlack,
			cell.ColorBlack,
			cell.ColorBlack,
		}),
		barchart.ShowValues(),
		barchart.BarWidth(8),
		barchart.Labels([]string{
			"🐮🐮 strong buy",
			"🐮 buy",
			"hold",
			"🐻 sell",
			"🐻🐻 strong sell",
		}),
	)
	if err != nil {
		panic(err)
	}
	go widgets.UpdateTrendBarChart(ctx, bcTarget, "target")

	// container outlay
	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("👋 PRESS Q TO QUIT"),
		container.SplitVertical(
			container.Left(
				container.SplitHorizontal(
					container.Top(
						container.SplitHorizontal(
							container.Top(
								container.PlaceWidget(timeNow),
							),
							container.Bottom(
								container.SplitHorizontal(
									container.Top(
										container.Border(linestyle.Light),
										container.BorderTitle("📈 stock prices by yahoo finance "),
										container.PlaceWidget(stk),
									),
									container.Bottom(
										container.SplitHorizontal(
											container.Top(
												container.Border(linestyle.Light),
												container.BorderTitle(fmt.Sprintf("%s stock price by yahoo finance", cfg.Stocks[0])),
												container.PlaceWidget(lc),
											),
											container.Bottom(
												container.PlaceWidget(volchart),
											),
											container.SplitPercent(70),
										),
									),
									container.SplitPercent(30),
								),
							),
							container.SplitPercent(5),
						),
					),
					container.Bottom(
						container.SplitHorizontal(
							container.Top(
								container.Border(linestyle.Light),
								container.BorderTitle("📅 earning calendars by finnhub.io "),
								container.PlaceWidget(earningCals),
							),
							container.Bottom(
								container.SplitHorizontal(
									container.Top(
										container.Border(linestyle.Light),
										container.BorderTitle("📨 company news by finnhub.io "),
										container.PlaceWidget(companyNewsRoll),
									),
									container.Bottom(
										container.Border(linestyle.Light),
										container.BorderTitle("💬 market news by finnhub.io "),
										container.PlaceWidget(news),
									),
								),
							),
							container.SplitPercent(30),
						),
					),
					container.SplitPercent(45),
				),
			),
			container.Right(
				container.SplitHorizontal(
					container.Top(
						//container.Border(linestyle.Light),
						container.SplitHorizontal(
							container.Top(
								container.SplitHorizontal(
									container.Top(
										container.Border(linestyle.Light),
										container.BorderTitle("📂 basic financials by finnhub.io "),
										container.PlaceWidget(basicFinancials),
									),
									container.Bottom(
										//container.Border(linestyle.Light), //"download as csv button"
										container.PlaceWidget(export1),
										container.AlignHorizontal(align.HorizontalRight),
									),
									container.SplitPercent(85),
								),
							),
							container.Bottom(
								container.SplitHorizontal(
									container.Top(
										container.Border(linestyle.Light),
										container.BorderTitle(fmt.Sprintf("🗂️ [%s] financials as reported by finnhub.io ", cfg.Stocks[0])),
										container.PlaceWidget(financialReports),
									),
									container.Bottom(
										//container.Border(linestyle.Light), //"download as csv button"
										container.PlaceWidget(export2),
										container.AlignHorizontal(align.HorizontalRight),
									),
									container.SplitPercent(85),
								),
							),
						),
					),
					container.Bottom(
						container.Border(linestyle.Light),
						container.BorderTitle("🗂️ stock estimates by finnhub.io "),
						container.SplitVertical(
							container.Left(
								container.SplitVertical(
									container.Left(
										container.Border(linestyle.Light),
									),
									container.Right(
										container.Border(linestyle.Light),
									),
								),
							),
							container.Right(
								container.SplitVertical(
									container.Left(
										container.Border(linestyle.Light),
									),
									container.Right(
										container.Border(linestyle.Light),
									),
								),
							),
						),
					),
					container.SplitPercent(75),
				),
			),
			container.SplitPercent(35),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// define the quitter
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		log.Fatal(err)
	}
}
