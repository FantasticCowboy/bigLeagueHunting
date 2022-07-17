package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationController struct {
	frameWidth      int
	frameHeight     int
	totalFrameCount int
	currentFrame    int
	frames          *ebiten.Image
}

func CreateAnimationController(
	frameWidth int,
	frameHeight int,
	totalFrameCount int,
	frames *ebiten.Image) *AnimationController {
	return &AnimationController{
		frameWidth:      frameWidth,
		frameHeight:     frameHeight,
		totalFrameCount: totalFrameCount,
		frames:          frames,
		currentFrame:    0,
	}
}

func (controller *AnimationController) getFrameBox(frameNumber int) *image.Rectangle {
	return &image.Rectangle{
		image.Pt(controller.frameWidth*controller.currentFrame, 0),
		image.Pt(controller.frameWidth*(controller.currentFrame+1), controller.frameHeight),
	}
}

func (controller *AnimationController) GetFrame() *ebiten.Image {
	frameNumber := controller.currentFrame

	controller.currentFrame++

	if controller.currentFrame >= controller.totalFrameCount {
		controller.currentFrame = 0
	}

	return ebiten.NewImageFromImage(
		controller.frames.SubImage(
			*controller.getFrameBox(frameNumber),
		))
}
