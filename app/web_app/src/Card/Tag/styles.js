// Tag stylesheet

import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
    tag: {
        paddingVertical: 8,
        paddingHorizontal: 12,
        marginRight: 10,
        marginBottom: 10,
        textAlign: 'center',
        borderRadius: 15,
        borderWidth: 2,
        fontSize: 15,
        letterSpacing: 0.2,
    },

    tagMatch: {
        backgroundColor: '#c8f1c8',
        borderColor: "#77dd77",
        color: '#77dd77'
    },

    tagNoMatch: {
        backgroundColor: '#ffb2ae',
        borderColor: '#ff6961',
        color: '#ff6961'
    }
}); 