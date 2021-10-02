import React from "react";
import { View, ScrollView, Text } from 'react-native';
import { styles } from './styles';

export default function BottomSheet({job}) {
    return (
        <View style={styles.jobDescription}>
            <View>
                <Text>Description</Text>
            </View>

            <ScrollView showsHorizontalScrollIndicator={false}>
                
                <Text>{job.description}</Text>
            </ScrollView>
        </View>
    );
}