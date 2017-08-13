package appui

import (
	"testing"

	"github.com/moncho/dry/docker"
	"github.com/moncho/dry/mocks"
	"github.com/moncho/dry/ui"
)

func TestImagesToShowSmallScreen(t *testing.T) {
	_ = "breakpoint"
	daemon := &mocks.DockerDaemonMock{}
	imagesLen := daemon.ImagesCount()
	if imagesLen != 5 {
		t.Errorf("Daemon has %d images, expected %d", imagesLen, 3)
	}

	cursor := ui.NewCursor()
	ui.ActiveScreen = &ui.Screen{
		Dimensions: &ui.Dimensions{Height: 15, Width: 100},
		Cursor:     cursor}

	renderer := NewDockerImagesWidget(0)
	imagesFromDaemon, _ := daemon.Images()
	renderer.PrepareToRender(NewDockerImageRenderData(
		imagesFromDaemon, docker.NoSortImages))

	images := renderer.visibleRows()
	if len(images) != 4 {
		t.Errorf("Images renderer is showing %d images, expected %d", len(images), 4)
	}
	if images[0].ID.Text != "8dfafdbc3a40" {
		t.Errorf("First image rendered is %s, expected %s. Cursor: %d", images[0].ID.Text, "8dfafdbc3a40", cursor.Position())
	}

	if images[2].ID.Text != "26380e1ca356" {
		t.Errorf("Last image rendered is %s, expected %s. Cursor: %d", images[2].ID.Text, "26380e1ca356", cursor.Position())
	}
	cursor.ScrollTo(4)
	renderer.PrepareToRender(NewDockerImageRenderData(
		imagesFromDaemon, docker.NoSortImages))
	images = renderer.visibleRows()
	if len(images) != 4 {
		t.Errorf("Images renderer is showing %d images, expected %d", len(images), 4)
	}
	if images[0].ID.Text != "541a0f4efc6f" {
		t.Errorf("First image rendered is %s, expected %s. Cursor: %d", images[0].ID.Text, "541a0f4efc6f", cursor.Position())
	}

	if images[2].ID.Text != "a3d6e836e86a" {
		t.Errorf("Last image rendered is %s, expected %s. Cursor: %d", images[2].ID.Text, "a3d6e836e86a", cursor.Position())
	}
}

func TestImagesToShow(t *testing.T) {
	_ = "breakpoint"
	daemon := &mocks.DockerDaemonMock{}
	imagesLen := daemon.ImagesCount()
	if imagesLen != 5 {
		t.Errorf("Daemon has %d images, expected %d", imagesLen, 3)
	}

	cursor := ui.NewCursor()

	ui.ActiveScreen = &ui.Screen{Dimensions: &ui.Dimensions{Height: 20, Width: 100},
		Cursor: cursor}
	renderer := NewDockerImagesWidget(0)

	imagesFromDaemon, _ := daemon.Images()
	renderer.PrepareToRender(NewDockerImageRenderData(
		imagesFromDaemon, docker.NoSortImages))

	images := renderer.visibleRows()
	if len(images) != 5 {
		t.Errorf("Images renderer is showing %d images, expected %d", len(images), 5)
	}
	cursor.ScrollTo(3)
	images = renderer.visibleRows()
	if len(images) != 5 {
		t.Errorf("Images renderer is showing %d images, expected %d", len(images), 5)
	}
	if images[0].ID.Text != "8dfafdbc3a40" {
		t.Errorf("First image rendered is %s, expected %s", images[0].ID.Text, "8dfafdbc3a40")
	}

	if images[4].ID.Text != "03b4557ad7b9" {
		t.Errorf("Last image rendered is %s, expected %s", images[4].ID.Text, "03b4557ad7b9")
	}
}

func TestImagesToShowNoImages(t *testing.T) {
	renderer := NewDockerImagesWidget(0)

	renderer.PrepareToRender(NewDockerImageRenderData(
		nil, docker.NoSortImages))

	images := renderer.visibleRows()
	if len(images) != 0 {
		t.Error("Unexpected number of image rows, it should be 0")
	}

}
