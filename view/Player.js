export default class Player {
    constructor(offsetX, offsetY) {
        this.offset = {
            x: offsetX,
            y: offsetY
        }
        this.width = offsetX / 10;
        this.height = offsetY / 10;
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
    }
}