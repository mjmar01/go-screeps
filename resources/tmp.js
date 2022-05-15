let c1 = Game.creeps["Frank"]
let c2 = Game.creeps["Jeff"]
let s = Game.getObjectById("ef990774d80108c");
let site = Game.getObjectById("5ce39cf83979762");

if (Game.time % 2 === 0) {
    c1.harvest(s)
    c2.harvest(s)
} else {
    c1.build(site)
    c2.build(site)
}