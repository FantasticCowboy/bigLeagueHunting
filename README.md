# ebitenTutorial

This game will help me get familar with ebiten and hone my skills in game development.

I want this game to be similar to duck shooter https://softfamous.com/wp-content/uploads/2018/10/Duck-Shooter.jpg.

The first iteration of this game will be red square will appear on the screen slowly. The player will then shoot the square and gain points. The concept art can be found here. https://docs.google.com/presentation/d/1uJ5Ls6_tUSrG_hbFxt3fMM18Wc2cxvzXtAcdGeCwass/edit#slide=id.p

For this initial version, there will be three states for the game - title, playing, and game over.

The title state can transition into the playing. The playing state is the player actually playing the game. Once the game ends, the game enters into the game over state. Once in the game over state, the player will have an option to restart the game.

For this intitial version to run properly, the following objects and their respective logic needs to be implemented.

- Game objects
  - Cursor AKA reticle AKA gun
  - Duck
  - Scoreboard
- Menu Objects
  - Title
  - Pause? - do this later on
  - Game Over

Because golang uses composition instead of inheritance, I should make the following sub-classes.
  - Sprite
    - 2d image
    - Translate/Rotation/Resizing
