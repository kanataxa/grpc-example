import Player from './Player.js';
import GameLoop from './GameLoop.js';
const renderer = PIXI.autoDetectRenderer(660, 660, {backgroundColor: 0x34495e, antialias: true});
document.body.appendChild(renderer.view);

//const {HelloRequest, HelloReply} = require('./server_pb.js');
//const {GreeterClient} = require('./server_grpc_web_pb.js');

const player = new Player(renderer.width, renderer.height);

const stage = new PIXI.Container();
stage.addChild(player.box);

const render = () => {
    renderer.render(stage)
}
const updater = dt => {
    player.update(dt)
}
const loop = new GameLoop({renderer:render, updater});
loop.start()
