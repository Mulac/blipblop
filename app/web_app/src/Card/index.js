import React, { useRef, useMemo } from "react";
import { View, Text, Animated, Image, ScrollView, PanResponder } from "react-native";
import { styles } from "./styles";
import BottomSheet, {BottomSheetScrollView} from '@gorhom/bottom-sheet';


export default function Card({ job, isFirst, swipe, ...rest }) {
    const rotate = swipe.x.interpolate({
        inputRange: [-100, 0, 100],
        outputRange: ['8deg', '0deg', '-8deg'],
    });

    const animatedCardSwipe = {
        transform: [...swipe.getTranslateTransform(), {rotate}]
    };
    
    const snapPoints = useMemo(() => ['30%', '90%'], []);


    return (
        <Animated.View style={[styles.container, isFirst && animatedCardSwipe]} {...rest}>
            <View style={styles.jobDetails}>
                <View style={styles.companyDetails}>
                    { job.CompanyImage !== "" && 
                        <Image style={styles.companyImage}
                            source={
                                {uri: job.CompanyImage}
                            }
                        /> 
                    }
                    <Text style={styles.companyName}>{job.Company}</Text>
                </View>

                <Text style={styles.positionName}>{job.Title}</Text>
                
                <View style={styles.locationDetails}>
                    <Image
                        style={styles.locationPin}
                        source={require('../../assets/pin.png')} 
                    />  
                    <Text>{job.Location}</Text>
                </View>
            </View>
            
            <BottomSheet
                index={0}
                snapPoints={snapPoints}
                enableOverDrag={false}
            >
                
                <View style={styles.bottomSheet}>
                    <Text style={styles.descriptionText}>Description</Text>
                </View>
                <BottomSheetScrollView showsHorizontalScrollIndicator={false} style={styles.jobDescription}>         
                    <Text style={styles.jobDescriptionText}>{job.Description}</Text>
                </BottomSheetScrollView>
            </BottomSheet>
        </Animated.View>
    );
}