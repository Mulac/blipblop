import React from "react";
import { View, Text, Animated } from "react-native";
import { styles } from "./styles";

export default function Card({ job, isFirst, swipe, ...rest }) {

    const rotate = swipe.x.interpolate({
        inputRange: [-100, 0, 100],
        outputRange: ['8deg', '0deg', '-8deg'],
    });

    const animatedCardSwipe = {
        transform: [...swipe.getTranslateTransform(), {rotate}]
    };

    return (
        <Animated.View style={[styles.container, isFirst && animatedCardSwipe]} {...rest}>
            <Text style={styles.positionName}>Job Title: {job.positionName}</Text>
            {/* <Text>some more stuff to go here...</Text> */}
            {isFirst}
        </Animated.View>
    );
}