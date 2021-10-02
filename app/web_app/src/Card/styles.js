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

        borderTopLeftRadius: 30,
        borderTopRightRadius: 30,
        position: 'absolute',
        flex: 1,
        bottom: 0
    },
    containerFirst: {
        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 2,
        },
        shadowOpacity: 0.5,
        shadowRadius: 5,
        elevation: 5,
    },
    companyDetails: {
        flex: 1,
        marginHorizontal: 20,
        top: 20,
    },
    positionName: {
        flex: 1,
        fontSize: 20,
        fontWeight: 'bold',
        color: 'black'
    },
    companyImage: {
        width: 70,
        height: 50,

        // TODO(sam): Need to find a way to stop this flickering on re-rending
        resizeMode: 'contain'
    }
});