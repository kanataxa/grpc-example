
const renderer = PIXI.autoDetectRenderer(660, 660, {backgroundColor: 0x34495e, antialias: true});
document.body.appendChild(renderer.view);


class Player {
    constructor() {
        this.width = renderer.width / 10;
        this.height = renderer.height / 10;
        this.box = new PIXI.Graphics();
        this.init();
        document.addEventListener('keydown', key => { this.kewDown(key)});
    }
    init() {
        this.box.beginFill(0x3498db); // Blue color
        this.box.drawRect(0, 0, this.width, this.height);
        this.box.endFill();
    }

    kewDown(key) {

        if (key.keyCode === 87 || key.keyCode === 38) {
            if (this.box.position.y != 0) {
                this.box.position.y -= this.height;
            }
        }

        if (key.keyCode === 83 || key.keyCode === 40) {
            if (this.box.position.y != renderer.height - this.height) {
                this.box.position.y += this.height;
            }
        }

        if (key.keyCode === 65 || key.keyCode === 37) {
            if (this.box.position.x != 0) {
                this.box.position.x -= this.width;
            }
        }

        if (key.keyCode === 68 || key.keyCode === 39) {
            if (this.box.position.x != renderer.width - this.width) {
                this.box.position.x += this.width;
            }
        }
    }
}

class GameLoop {
    constructor() {
        this.duration = 0;
        this.rate = 30.0;
        this.tickInterval = 1000 / this.rate; // 0.016666s = 16.66666ms

        this.lastTick = 0;
        console.log(this)

    }

    start(renderer) {
        const current = performance.now();
        if (this.lastTick === 0) {
            this.lastTick = current
        }
        this.duration += current - this.lastTick;
        this.lastTick = current;
        if (this.duration < this.tickInterval) {
            requestAnimationFrame(() => {this.start(renderer)});
            return;
        }
        while (this.duration > this.tickInterval) {
            // game state update()
            this.duration -= this.tickInterval
        }
        renderer.render(stage);
        requestAnimationFrame(() => {this.start(renderer)});
    }

}

const player = new Player();

const stage = new PIXI.Container();
stage.addChild(player.box);

const loop = new GameLoop();
loop.start(renderer)
