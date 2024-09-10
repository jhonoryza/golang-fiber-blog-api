export const getCSRFToken = () => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; csrf_=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
};

export function mergeURLSearchParams(params, additionalParams) {
    const mergedParams = new URLSearchParams(params);

    additionalParams.forEach((value, key) => {
        mergedParams.set(key, value);
    });

    return mergedParams;
}

export function urlSearchParamsToObject(urlSearchParams) {
    const paramsObject = {};
    urlSearchParams.forEach((value, key) => {
        paramsObject[key] = value;
    });
    return paramsObject;
}

export function objectToURLSearchParams(obj, parentKey = '') {
    const params = new URLSearchParams();

    for (const key in obj) {
        if (obj.hasOwnProperty(key)) {
            const value = obj[key];
            const newKey = parentKey ? `${parentKey}[${key}]` : key;

            if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
                const nestedParams = objectToURLSearchParams(value, newKey);
                nestedParams.forEach((nestedValue, nestedKey) => {
                    params.append(nestedKey, nestedValue);
                });
            } else {
                params.append(newKey, value);
            }
        }
    }

    return params;
}