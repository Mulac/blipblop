import React, { useCallback, useEffect, useRef, useState } from 'react';
import { View, Animated, PanResponder } from 'react-native';
import Card from '../Card';
import { styles } from './styles';
import api from '../utils/api';

const axios = require('axios');

export default function Main() {
    const [jobs, setJobs] = useState([]);
    const swipe = useRef(new Animated.ValueXY()).current;

    // We change the DOM which counts as a side effect,
    // so this hook is ran every time we swipe
    useEffect(() => {
        if (jobs.length <= 1) {    
            
            axios.get('http://178.79.148.75/jobs', {
                auth: {
                    username: 'admin',
                    password: 'blipblop'
                }
            })
            .then((response) => {
                if (response.data != undefined) {
                    setJobs(jobs.concat(response.data));
                }
            })
            .catch((error) => {
                console.log(error);
                setJobs(jobs.concat([]));
            });
        }
    }, [jobs.length]);

    const panResponder = PanResponder.create({
        onMoveShouldSetPanResponder: (_, { dx, dy}) => {
            const draggedLeft = dx < -15;
            const draggedRight = dx > 15;

            if (draggedLeft || draggedRight)
                return true;
        },
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
                        x: direction * 700,
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

                // We only want to handle draggin on the first card
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
