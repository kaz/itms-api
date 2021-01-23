package itms

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type (
	Action func(browser.Submittable) (bool, error)

	Interactor struct {
		name  string
		index int
		acts  []Action
	}
)

func interact(name string, index int) *Interactor {
	return &Interactor{name, index, []Action{}}
}
func (i *Interactor) withInput(inputs map[string]string) *Interactor {
	i.acts = append(i.acts, func(form browser.Submittable) (bool, error) {
		for key, value := range inputs {
			if err := form.Input(key, value); err != nil {
				return false, fmt.Errorf("input %s: %w", key, err)
			}
		}
		return true, nil
	})
	return i
}
func (i *Interactor) whenActionIs(action string) *Interactor {
	i.acts = append(i.acts, func(form browser.Submittable) (bool, error) {
		return strings.Contains(form.Dom().Find("[type='submit']").AttrOr("value", ""), action), nil
	})
	return i

}

func (i *Interactor) run(browser *browser.Browser) error {
	form, err := browser.Form(fmt.Sprintf("[name='%s']:nth-of-type(%d)", i.name, i.index))
	if err != nil {
		return fmt.Errorf("find: %w", err)
	}

	for _, act := range i.acts {
		ok, err := act(form)
		if err != nil {
			return fmt.Errorf("action: %w", err)
		}
		if !ok {
			fmt.Println("Skipped")
			return nil
		}
	}

	if err := form.Submit(); err != nil {
		return fmt.Errorf("submit: %w", err)
	}
	if browser.StatusCode() != http.StatusOK {
		return fmt.Errorf("http: code=%d", browser.StatusCode())
	}

	return nil
}

func jemaControl(action string, index int) error {
	browser := surf.NewBrowser()

	err := browser.Open("https://itms.hitsecurity.jp")
	if err != nil {
		return fmt.Errorf("open hitsecurity: %w", err)
	}

	if err := interact("IMLG02LoginBean", 1).run(browser); err != nil {
		return fmt.Errorf("submit [index]: %w", err)
	}
	if err := interact("IMLG02LoginBean", 1).withInput(credential).run(browser); err != nil {
		return fmt.Errorf("submit form IMLG02LoginBean: %w", err)
	}
	if err := interact("IMEO02EnergyBean", 1).run(browser); err != nil {
		return fmt.Errorf("submit form IMEO02EnergyBean: %w", err)
	}
	if err := interact("IMEO06RemoteBean", 1).run(browser); err != nil {
		return fmt.Errorf("submit form IMEO06RemoteBean: %w", err)
	}
	if err := interact("IMEO07JEMAControlBean", index).run(browser); err != nil {
		return fmt.Errorf("submit form IMEO07JEMAControlBean: %w", err)
	}
	if err := interact("IMEO07JEMAControlBean", 1).whenActionIs(action).run(browser); err != nil {
		return fmt.Errorf("submit [target]: %w", err)
	}

	return nil
}
