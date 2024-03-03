package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	INDAH_WATER_URL = "https://customerportal.iwk.com.my/pay_bill"
)

type IndahWaterAccountInfo struct {
	AccountNumber string
	Address       string
	BillAmount    string
	BillReference string
}

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-notifications", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	data := new(IndahWaterAccountInfo)

	// run task list
	err := chromedp.Run(ctx,
		chromedp.Navigate(INDAH_WATER_URL),
		chromedp.WaitVisible(`#fmCheckBal`, chromedp.ByID),
		// chromedp.WaitVisible(`body form`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Focus(`#cb-san`, chromedp.ByID),
		chromedp.SendKeys(`#cb-san`, "89753172", chromedp.ByID),
		chromedp.Sleep(1*time.Second),
		chromedp.Click(`#btnPaySubmit`, chromedp.ByID),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(`div.form-control-txt.san`, &data.AccountNumber, chromedp.ByQuery),
		chromedp.Text(`div.form-control-txt.address`, &data.Address, chromedp.ByQuery),
		chromedp.Value(`#pay-amount`, &data.BillAmount, chromedp.ByID),
		chromedp.Text(`div.form-control-txt.bill_no`, &data.BillReference, chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account Number: ", data.AccountNumber)
	fmt.Println("Address: ", data.Address)
	fmt.Println("Bill Amount: ", data.BillAmount)
	fmt.Println("Bill Reference: ", data.BillReference)

}
