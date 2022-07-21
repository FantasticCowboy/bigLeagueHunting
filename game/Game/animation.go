package game

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationController struct {
	FrameWidth      int
	FrameHeight     int
	TotalFrameCount int
	CurrentFrame    int
	FramesPerSecond int
	Frames          *ebiten.Image
	lastUpdate      time.Time
}

func CreateAnimationController(
	frameWidth int,
	frameHeight int,
	totalFrameCount int,
	frames *ebiten.Image,
	fps int,
) *AnimationController {

	return &AnimationController{
		FrameWidth:      frameWidth,
		FrameHeight:     frameHeight,
		TotalFrameCount: totalFrameCount,
		Frames:          frames,
		CurrentFrame:    0,
		FramesPerSecond: fps,
		lastUpdate:      time.Now(),
	}
}

func (controller *AnimationController) getFrameBox(frameNumber int) *image.Rectangle {
	return &image.Rectangle{
		image.Pt(controller.FrameWidth*controller.CurrentFrame, 0),
		image.Pt(controller.FrameWidth*(controller.CurrentFrame+1), controller.FrameHeight),
	}
}

func (controller *AnimationController) GetFrame() *ebiten.Image {
	frameNumber := controller.CurrentFrame

	if time.Since(controller.lastUpdate) > (time.Second / time.Duration(controller.FramesPerSecond)) {
		controller.CurrentFrame++
		controller.lastUpdate = time.Now()
	}

	if controller.CurrentFrame >= controller.TotalFrameCount {
		controller.CurrentFrame = 0
	}

	return ebiten.NewImageFromImage(
		controller.Frames.SubImage(
			*controller.getFrameBox(frameNumber),
		))
}
