from main import dp
from aiogram import types


@dp.message_handler(commands=['start'])
async def send_welcome(message: types.Message):
    await message.answer(f"Welcome to English words test")


