package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Action interface {
	Call(arg any)
}

type Event struct {
	subscribers []Action
}

func (e Event) Invoke(in any) {
	size := len(e.subscribers)
	for i := 0; i < size; i++ {
		e.subscribers[i].Call(in)
	}
}

func (e *Event) Add(action Action) {
	(*e).subscribers = append((*e).subscribers, action)
}

func (e *Event) Remove(action Action) {
	size := len((*e).subscribers)
	for i := 0; i < size; i++ {
		if action == (*e).subscribers[i] {
			(*e).subscribers[i] = (*e).subscribers[size-1]
			size -= 1
		}
	}

	(*e).subscribers = (*e).subscribers[:size]
}

type Elm struct {
	lines []string
}

func (e *Elm) Add(s string) {
	(*e).lines = append((*e).lines, s)
}

type Vertex struct {
	s string
}

func (v *Vertex) Call(arg any) {

}

func main() {
	/*
		mySlice := []int{}
		fmt.Println(mySlice)

		for i := 0; i < size; i++ {
			mySlice = append(mySlice, 1)
		}
		fmt.Println(mySlice[123])

		// FOR
		startTime := time.Now()
		var sum int = 0
		for i := 0; i < size; i++ {
			sum++
		}
		if sum != size {
			fmt.Printf("error: %d, %d", size, sum)
		}
		diff := time.Since(startTime)
		var timePassedFOR = diff

		// FOREACH
		startTime = time.Now()
		sum = 0
		for _, _ = range mySlice {
			sum++
		}
		if sum != size {
			fmt.Printf("error: %d, %d", size, sum)
		}
		diff = time.Since(startTime)
		var timePassedFOREACH = diff
		fmt.Println()
		fmt.Printf("END: %d, %d", size, sum)

		println()
		fmt.Printf("FOR: %d\n", timePassedFOR)
		fmt.Printf("FOREACH: %d", timePassedFOREACH)
	*/
	/*
		var slice []string
		RecursivePassList(&slice, 5)
		fmt.Println(slice)

		a := app.New()

		w := a.NewWindow("Misra Checker")
		w.Resize(fyne.NewSize(640, 320))
		var scene *fyne.Container = container.NewVBox()

		projectRoot := widget.NewEntry()
		projectRoot.SetPlaceHolder("projectRoot...")
		scene.Add(projectRoot)

		projectTestDir := widget.NewEntry()
		projectRoot.SetPlaceHolder("projectTestDir...")
		scene.Add(projectTestDir)

		reportDir := widget.NewEntry()
		projectRoot.SetPlaceHolder("reportDir...")
		scene.Add(reportDir)

		lintInstallationFolder := widget.NewEntry()
		projectRoot.SetPlaceHolder("lintInstallationFolder...")
		scene.Add(lintInstallationFolder)

		lintButton := widget.NewButton("Save", func() {
			log.Println("Content was:", projectRoot.Text)
		})
		scene.Add(lintButton)

		w.SetContent(scene)
		w.Show()

		a.Run()
		fmt.Println("Exiting")
	*/

	var vertex Vertex
	vertex.s = "ciao"

	var act Action = &vertex

	e := Event{}
	e.Add(act)

	e.Invoke(any("ciao"))
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func RecursivePassList(l *[]string, index int) *[]string {
	if index == 0 {
		*l = append(*l, "123")
		return l
	}

	index -= 1
	*l = append(*RecursivePassList(l, index), "123")
	return l
}

func CallBack() {
	fmt.Println("CALLBACK")
}
