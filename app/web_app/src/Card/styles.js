// Card stylesheet

import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
    container: {
        width: '100%',
        height: '90%',
        // backgroundColor: '#FF6961',
        backgroundColor: '#fff',
        borderTopWidth: 1,
        borderLeftWidth: 1,
        borderRightWidth: 1,
        borderColor: 'grey',

        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 2,
        },
        shadowOpacity: 0.25,
        shadowRadius: 3.84,
        elevation: 5,

        borderTopLeftRadius: 30,
        borderTopRightRadius: 30,
        position: 'absolute',
        flex: 1,
        alignItems: 'center',
        bottom: 0
    },
    positionName: {
        top: 50,
        paddingHorizontal: 20,
        fontSize: 20,
        fontWeight: 'bold',
        color: 'black'
    }
});