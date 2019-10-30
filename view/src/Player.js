import Backend from "./Backend.js";

class InputQueue {
    constructor() {
        this.maxLength = 20;
        this.inputs = []
    }
    add(key) {
        this.inputs[this.inputs.length % this.maxLength] = key;
    }
    pop() {
        if (this.inputs.length === 0) return null;
        return this.inputs.shift();
    }
}

const recoverFrame = {
    37: 5,
    38: 2,
    39: 10,
    40: 2,
    65: 2,
    68: 7,
    83: 5,
    87: 1,
};

export default class Player {
    constructor(offsetX, offsetY) {
        this.offset = {
            x: offsetX,
            y: offsetY
        };
        this.rigidTick = 0;
        this.inputQueue = new InputQueue();
        this.backend = new Backend();
        this.width = offsetX / 10;
        this.height = offsetY / 10;
        this.box = new PIXI.Graphics();
        this.init();
        document.addEventListener('keydown', key => {
            this.backend.add(key);
            this.inputQueue.add(key);
        })
    }
    init() {
        this.box.beginFill(0x3498db); // Blue color
        this.box.drawRect(0, 0, this.width, this.height);
        this.box.endFill();
    }

    update() {
        const key = this.inputQueue.pop()
        if (key && this.rigidTick === 0) {
            this.kewDown(key);
        } else if (this.rigidTick > 0) {
            this.rigidTick -= 1;
        }
    }

    kewDown(key) {
        if (key.keyCode === 87 || key.keyCode === 38) {
            if (this.box.position.y != 0) {
                this.box.position.y -= this.height;
            }
        }

        if (key.keyCode === 83 || key.keyCode === 40) {
            if (this.box.position.y != this.offset.y - this.height) {
                this.box.position.y += this.height;
            }
        }

        if (key.keyCode === 65 || key.keyCode === 37) {
            if (this.box.position.x != 0) {
                this.box.position.x -= this.width;
            }
        }

        if (key.keyCode === 68 || key.keyCode === 39) {
            if (this.box.position.x != this.offset.x - this.width) {
                this.box.position.x += this.width;
            }
        }
        this.rigidTick += recoverFrame[key.keyCode]
    }
}