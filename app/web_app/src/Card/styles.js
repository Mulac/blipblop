// Card stylesheet

import { StyleSheet } from "react-native";

export const styles = StyleSheet.create({
    container: {
        width: '100%',
        height: '90%',
        backgroundColor: '#fff',
        borderTopWidth: 1,
        borderLeftWidth: 1,
        borderRightWidth: 1,
        borderColor: 'grey',

        borderTopLeftRadius: 30,
        borderTopRightRadius: 30,
        position: 'absolute',
        flex: 1,
        bottom: 0,
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
        flexDirection: 'row',
        alignItems: 'center'
    },  
    companyName: {
        fontSize: 15,
        fontWeight: 'bold'
    },
    positionName: {
        marginTop: 10,
        fontSize: 20,
        fontWeight: 'bold',
        color: 'black'
    },
    companyImage: {
        width: 70,
        height: 50,
        marginRight: 10,

        // TODO(sam): Need to find a way to stop this flickering on re-rending
        resizeMode: 'contain'
    },

    locationDetails: {
        flexDirection: 'row',
        alignItems: 'center'
    },
    locationPin: {
        marginTop: 10,
        marginRight: 10,
        width: 30,
        height: 30,
    },

    jobDetails: {
        margin: 20
    },

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