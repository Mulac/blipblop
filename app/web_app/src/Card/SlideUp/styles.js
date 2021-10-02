// BottomSheet stylesheet

import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
    jobDescription: {
        // marginTop: '50%',
        backgroundColor: '#f2f2f2',
        paddingTop: 20,
        paddingHorizontal: 20,
    },

    bottomSheet: {
        backgroundColor: '#f2f2f2',
        paddingHorizontal: 20,
        paddingVertical: 20,
        borderTopLeftRadius: 20,
        borderTopRightRadius: 20,
        borderBottomWidth: 1,
        borderColor: 'grey'
    },

    description: {
        fontSize: 15,
        fontWeight: 'bold',
    }
});