import React, { useRef } from "react";
import { View, Text, Animated, Image, ScrollView, PanResponder } from "react-native";
import { styles } from "./styles";
import SlideUp from "./SlideUp";

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
            <View style={styles.jobDetails}>
                <View style={styles.companyDetails}>
                    { job.companyImage && 
                        <Image style={styles.companyImage}
                            source={
                                {uri: job.companyImage}
                            }
                        /> 
                    }
                    <Text style={styles.companyName}>{job.company}</Text>
                </View>

                <Text style={styles.positionName}>{job.positionName}</Text>
                
                <View style={styles.locationDetails}>
                    <Image
                        style={styles.locationPin}
                        source={require('../../assets/pin.png')} 
                    />  
                    <Text>{job.location}</Text>
                </View>
            </View>

            <SlideUp
                job={job}
            />
        </Animated.View>
    );
}