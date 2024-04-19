<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AM-QR</title>
    <link rel="stylesheet" href="./style/style.css">
</head>

<body class="scroll-smooth font-Poppins">
    <main class="text-lg border">
        <div class="border">
            <img src="./assets/icons/person_icon.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
            <input type="text" name="f_name" id="full_name" placeholder="Enter Your Full Name" class=" align-middle">
        </div>
        <div class="border">
            <input type="text" name="u_age" id="u_age" placeholder="Age">
        </div>
        <div class="border">
            <label>DOB:</label>
            <input type="date" name="u_dob" id="u_dob">
        </div>
        <div class="border">
            <img src="./assets/icons/mail_icon.svg" alt="mail" class="opacity-50 inline-block w-8 h-9 mr-2">
            <input type="email" name="u_mail" id="u_mail" placeholder="E-mail Id" class="align-middle">
        </div>
        <div class="border">
            <img src="./assets/icons/phone.svg" alt="phone" class="opacity-50 inline-block w-8 h-9 mr-2">
            <input type="number" name="u_num" id="u_num" placeholder="Phone Number" class="align-middle">
        </div>
        <div class="border">
            <img src="./assets/icons/institution.png" alt="institution" class=" opacity-50 inline-block mr-2 h-8 w-7">
            <input type="text" name="u_org" id="u_org" placeholder="Organisation" class="align-middle">
        </div>
        <div class="border">
            <img src="./assets/icons/username.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
            <input type="text" name="u_username" id="u_username" placeholder="Username" class="align-middle">
        </div>
        <div class="border">
            <img src="./assets/icons/password.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
            <input type="password" name="u_password" id="u_password" placeholder="Password" class="align-middle">
        </div>
        <div class="border">
            <label>Upload your Photo:
                <img src="./assets/icons/image_icon.svg" alt="image" class="opacity-50 inline-block w-8 h-9 mr-2">
            </label>
            <input type="file" name="u_img" id="u_img">
        </div>
    </main>
</body>

</html>