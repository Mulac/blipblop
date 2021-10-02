import React, { useCallback, useEffect, useRef, useState } from 'react';
import { View, Animated, PanResponder } from 'react-native';
import Card from '../Card';
import { jobs as jobsArray } from './data';
import { styles } from './styles';


export default function Main() {
    const [jobs, setJobs] = useState(jobsArray)
    const swipe = useRef(new Animated.ValueXY()).current;

    // We change the DOM which counts as a side effect,
    // so this hook is ran every time we swipe
    useEffect(() => {
        if (jobs.length <= 1) {
            setJobs(jobs.concat(jobsArray));
        }
    }, [jobs.length]);

    const panResponder = PanResponder.create({
        onMoveShouldSetPanResponder: () => true,
        // Move to the current x, y position of the gesture (finger on the screens location)
        onPanResponderMove: (_, { dx, dy }) => {
            swipe.setValue({ x: dx, y: 0 });
        },
        // When we release the swipe we want the card to bounce back to its start position
        onPanResponderRelease: (_, { dx, dy }) => {

            // remove the card
            const direction = Math.sign(dx);
            const isActionActive = Math.abs(dx) > 100;
            
            if (isActionActive) {
                Animated.timing(swipe, {
                    duration: 200,
                    toValue: {
                        x: direction * 500,
                        y: dy
                    },
                    useNativeDriver: true
                }).start(removeTopCard);

            } else {
                Animated.spring(swipe, {
                    // Start position
                    toValue: {
                        x: 0,
                        y: 0
                    },
                    useNativeDriver: true,
                    // Limiter for the speed we want to the card to bounce back to the start
                    friction: 7, 
                }).start();
            }
        }
    });

    const removeTopCard = useCallback(() => {
        setJobs((prevState) => prevState.slice(1));
        swipe.setValue({ x: 0, y: 0 });
    }, [swipe]);

    return (
        <View style={styles.container}>
            {jobs.map((jobObj, i) => {
                const isFirst = i === 0;
                const dragHandlers = isFirst ? panResponder.panHandlers : {};

                return (
                    <Card 
                        // We have to set a key for each card to react views them as unique
                        key={i}
                        job={jobObj}
                        isFirst={isFirst}
                        swipe={swipe}
                        {...dragHandlers}
                    />
                );
            }).reverse()}
        </View>
    );
}
