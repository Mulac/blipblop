import React from "react";
import { View, Text, Animated, Image } from "react-native";
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
        <Animated.View style={[styles.container, isFirst && animatedCardSwipe, isFirst && styles.containerFirst]} {...rest}>
            <View style={styles.companyDetails}>
                <View>
                    <Image
                        style={styles.companyImage}
                        source={
                            {uri: job.companyImage}
                        }
                    />
                </View>

                <Text style={styles.positionName}>{job.positionName}</Text>
            </View>
        </Animated.View>
    );
}