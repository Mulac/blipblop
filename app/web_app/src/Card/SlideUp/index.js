import React, { useMemo } from "react";
import { View, Text } from 'react-native';
import { styles } from './styles';
import BottomSheet, {BottomSheetScrollView} from '@gorhom/bottom-sheet';

export default function SlideUp({job}) {
    // variables
    const snapPoints = useMemo(() => ['30%', '60%'], []);

    // callbacks
    // const handleSheetChanges = useCallback((index) => {
    //     console.log('handleSheetChanges', index);
    // }, []);

    return (
       

        <BottomSheet
            index={0}
            snapPoints={snapPoints}
            // onChange={handleSheetChanges}
        >
            
            <View style={styles.bottomSheet}>
                <Text style={styles.descriptionText}>Description</Text>
            </View>
            <BottomSheetScrollView showsHorizontalScrollIndicator={false} style={styles.jobDescription}>         
                <Text>{job.description}</Text>
            </BottomSheetScrollView>
        </BottomSheet>
    );
}