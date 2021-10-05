const { Dimensions } = require('react-native');

const { width, height } = Dimensions.get('screen');

export const CARD = {
    dragBuffer: 0.2,
    friction: 7,
    outOfScreen: width + 0.5 * width,
    actionOffset: 100,
    outOfScreenDuration: 200
};