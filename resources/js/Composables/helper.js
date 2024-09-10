export const getCSRFToken = () => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; csrf_=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
};
