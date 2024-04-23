<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AM-QR</title>
    <link rel="stylesheet" href="./style/style.css">
</head>

<body class="scroll-smooth font-Poppins">
    <div class="w-screen h-screen max-w-full max-h-full flex flex-col md:flex-row justify-center items-center">
        <main class="bg-[#FBF9E4] rounded-2xl md:rounded-3xl md:flex items-center">
            <div class="w-96 h-96">
                <img src="./assets/logo/logo1024.png" alt="logo" class="inline-block max-w-full max-h-full md:rounded-full">
            </div>
            <form action="" method="post" class="md:px-20 md:py-16 md:rounded-3xl text-md">
                <div class="py-1">
                    <img src="./assets/icons/person_icon.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="text" name="f_name" id="full_name" placeholder="Enter Your Full Name" class=" align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/age.png" alt="age" class="opacity-50 inline-block w-8 h-8 mr-2">
                    <input type="text" name="u_age" id="u_age" placeholder="Age" class="bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <label>DOB:</label>
                    <input type="date" name="u_dob" id="u_dob" class="bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/mail_icon.svg" alt="mail" class="opacity-50 inline-block w-8 h-9 mr-2">
                    <input type="email" name="u_mail" id="u_mail" placeholder="E-mail Id" class="align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/phone.svg" alt="phone" class="opacity-50 inline-block w-8 h-9 mr-2">
                    <input type="text" name="u_num" id="u_num" placeholder="Phone Number" class="align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/institution.png" alt="institution" class=" opacity-50 inline-block mr-2 h-8 w-7">
                    <input type="text" name="u_org" id="u_org" placeholder="Organisation" class="align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/username.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="text" name="u_username" id="u_username" placeholder="Username" class="align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <img src="./assets/icons/password.svg" alt="person" class="opacity-50 inline-block w-8 h-10 mr-2">
                    <input type="password" name="u_password" id="u_password" placeholder="Password" class="align-middle bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-1">
                    <label>Upload your Photo:
                        <img src="./assets/icons/image_icon.svg" alt="image" class="opacity-50 inline-block w-8 h-9 mr-2">
                    </label>
                    <input type="file" name="u_img" id="u_img" class="bg-transparent placeholder:text-[#547781]">
                </div>
                <div class="py-2">
                    <input type="submit" class="w-[7rem] ml-4 bg-[#0b64f4] text-white px-4 py-2 rounded-md hover:bg-[#1d4fd7] active:bg-[bg-[#1d4fd7]] hover:scale-105   active:scale-110 transition-all ease-in-out delay-50 tracking-[2px]" value="Sign Up" name="submit">
                </div>
            </form>
        </main>
    </div>
</body>

</html>